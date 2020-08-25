package experimentstore

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally"
	"go.uber.org/zap"

	experimentation "github.com/lyft/clutch/backend/api/chaos/experimentation/v1"
	"github.com/lyft/clutch/backend/id"
	"github.com/lyft/clutch/backend/service"
	pgservice "github.com/lyft/clutch/backend/service/db/postgres"
)

const Name = "clutch.service.chaos.experimentation.store"

// ExperimentStore stores experiment data
type ExperimentStore interface {
	CreateExperiments(context.Context, []*experimentation.Experiment) error
	DeleteExperiments(context.Context, []uint64) error
	GetExperiments(context.Context) ([]*experimentation.Experiment, error)
	Close()
}

type experimentStore struct {
	db *sql.DB
}

// New returns a new NewExperimentStore instance.
func New(_ *any.Any, _ *zap.Logger, _ tally.Scope) (service.Service, error) {
	p, ok := service.Registry[pgservice.Name]
	if !ok {
		return nil, errors.New("could not find database service")
	}

	client, ok := p.(pgservice.Client)
	if !ok {
		return nil, errors.New("experiment store wrong type")
	}

	return &experimentStore{
		client.DB(),
	}, nil
}

func (fs *experimentStore) CreateExperiments(ctx context.Context, experiments []*experimentation.Experiment) error {
	// This API call will eventually be broken into 2 separate calls:
	// 1) creating the config
	// 2) starting a new experiment with the config

	// All experiments are created in a single transaction
	tx, err := fs.db.Begin()
	if err != nil {
		return err
	}

	for _, experiment := range experiments {
		// Step 1) create the config
		configID := id.NewID()

		configJson, err := marshalConfig(experiment)
		if err != nil {
			return err
		}

		configSql := `INSERT INTO experiment_config (id, details) VALUES ($1, $2)`
		_, err = fs.db.ExecContext(ctx, configSql, configID, configJson)
		if err != nil {
			return err
		}

		// Step 2) start a new experiment with the config
		runSql := `
			INSERT INTO experiment_run (
                id, 
		        experiment_config_id,
		        start_time,
			    end_time,                        
		        creation_time)
            VALUES ($1, $2, NOW(), NULL, NOW())`

		_, err = fs.db.ExecContext(ctx, runSql, id.NewID(), configID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}

// DeleteExperiments deletes the specified experiments from the store.
func (fs *experimentStore) DeleteExperiments(ctx context.Context, ids []uint64) error {
	sql := `DELETE FROM experiment_run WHERE id IN $1`
	_, err := fs.db.ExecContext(ctx, sql, ids)
	return err
}

// GetExperiments gets all experiments
func (fs *experimentStore) GetExperiments(ctx context.Context) ([]*experimentation.Experiment, error) {
	sql := `
        SELECT experiment_run.id, details FROM experiment_config, experiment_run
        WHERE experiment_config.id = experiment_run.experiment_config_id
        	AND experiment_run.start_time < now()`

	rows, err := fs.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var experiments []*experimentation.Experiment
	for rows.Next() {
		var experiment experimentation.Experiment
		var details string

		err = rows.Scan(&experiment.Id, &details)
		if err != nil {
			return nil, err
		}

		anyConfig := &any.Any{}
		if nil != jsonpb.Unmarshal(strings.NewReader(details), anyConfig) {
			return nil, err
		}

		experiment.Config = anyConfig
		experiments = append(experiments, &experiment)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return experiments, nil
}

// Close closes all resources held.
func (fs *experimentStore) Close() {
	fs.db.Close()
}

func marshalConfig(experiment *experimentation.Experiment) (string, error) {
	marshaler := jsonpb.Marshaler{}
	buf := &bytes.Buffer{}
	err := marshaler.Marshal(buf, experiment.GetConfig())
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

