DROP TABLE IF EXISTS experiments;
CREATE TABLE IF NOT EXISTS experiment_config (
    id BIGINT,
    details JSONB
);
CREATE TABLE IF NOT EXISTS experiment_run (
    id BIGINT,
    experiment_config_id BIGINT,
    start_time TIMESTAMP WITH TIME ZONE,
    end_time TIMESTAMP WITH TIME ZONE,
    scheduled_end_time TIMESTAMP WITH TIME ZONE,
    creation_time TIMESTAMP WITH TIME ZONE,
    termination_reason varchar(32)
);
