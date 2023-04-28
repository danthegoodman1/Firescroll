package utils

import "os"

var (
	Env_SleepSeconds           = MustEnvOrDefaultInt64("SHUTDOWN_SLEEP_SEC", 0)
	Env_ShutdownTimeoutSeconds = MustEnvOrDefaultInt64("SHUTDOWN_TIMEOUT_SEC", 1)

	Env_Namespace        = MustEnv("NAMESPACE")
	Env_ReplicaGroupName = MustEnv("REPLICA_GROUP")
	Env_InstanceID       = MustEnv("INSTANCE_ID")
	Env_KafkaSessionMs   = MustEnvOrDefaultInt64("KAFKA_SESSION_MS", 60_000)
	Env_NumPartitions    = MustEnvInt64("PARTITIONS")
	Env_TopicRetentionMS = MustEnvInt64("TOPIC_RETENTION_MS")

	Env_GCIntervalMs = MustEnvOrDefaultInt64("GC_INTERVAL_MS", 60_000*5) // 5 minute default
	Env_DBPath       = EnvOrDefault("DB_PATH", "/var/firescroll/db")

	Env_APIPort = EnvOrDefault("API_PORT", "8070")

	Env_Debug       = os.Getenv("DEBUG") == "1"
	Env_BadgerDebug = os.Getenv("BADGER_DEBUG") == "1"
)
