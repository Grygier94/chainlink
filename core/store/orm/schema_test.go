package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigSchema(t *testing.T) {
	items := map[string]string{
		"AdminCredentialsFile":                      "ADMIN_CREDENTIALS_FILE",
		"AllowOrigins":                              "ALLOW_ORIGINS",
		"AuthenticatedRateLimit":                    "AUTHENTICATED_RATE_LIMIT",
		"AuthenticatedRateLimitPeriod":              "AUTHENTICATED_RATE_LIMIT_PERIOD",
		"BalanceMonitorEnabled":                     "BALANCE_MONITOR_ENABLED",
		"BlockBackfillDepth":                        "BLOCK_BACKFILL_DEPTH",
		"BridgeResponseURL":                         "BRIDGE_RESPONSE_URL",
		"ChainID":                                   "ETH_CHAIN_ID",
		"ClientNodeURL":                             "CLIENT_NODE_URL",
		"DatabaseBackupFrequency":                   "DATABASE_BACKUP_FREQUENCY",
		"DatabaseBackupMode":                        "DATABASE_BACKUP_MODE",
		"DatabaseBackupURL":                         "DATABASE_BACKUP_URL",
		"DatabaseBackupDir":                         "DATABASE_BACKUP_DIR",
		"DatabaseListenerMaxReconnectDuration":      "DATABASE_LISTENER_MAX_RECONNECT_DURATION",
		"DatabaseListenerMinReconnectInterval":      "DATABASE_LISTENER_MIN_RECONNECT_INTERVAL",
		"DatabaseMaximumTxDuration":                 "DATABASE_MAXIMUM_TX_DURATION",
		"DatabaseTimeout":                           "DATABASE_TIMEOUT",
		"DatabaseURL":                               "DATABASE_URL",
		"DefaultHTTPAllowUnrestrictedNetworkAccess": "DEFAULT_HTTP_ALLOW_UNRESTRICTED_NETWORK_ACCESS",
		"DefaultHTTPLimit":                          "DEFAULT_HTTP_LIMIT",
		"DefaultHTTPTimeout":                        "DEFAULT_HTTP_TIMEOUT",
		"DefaultMaxHTTPAttempts":                    "MAX_HTTP_ATTEMPTS",
		"Dev":                                       "CHAINLINK_DEV",
		"EnableExperimentalAdapters":                "ENABLE_EXPERIMENTAL_ADAPTERS",
		"EnableLegacyJobPipeline":                   "ENABLE_LEGACY_JOB_PIPELINE",
		"EthBalanceMonitorBlockDelay":               "ETH_BALANCE_MONITOR_BLOCK_DELAY",
		"EthFinalityDepth":                          "ETH_FINALITY_DEPTH",
		"EthGasBumpPercent":                         "ETH_GAS_BUMP_PERCENT",
		"EthGasBumpThreshold":                       "ETH_GAS_BUMP_THRESHOLD",
		"EthGasBumpTxDepth":                         "ETH_GAS_BUMP_TX_DEPTH",
		"EthGasBumpWei":                             "ETH_GAS_BUMP_WEI",
		"EthGasLimitDefault":                        "ETH_GAS_LIMIT_DEFAULT",
		"EthGasLimitTransfer":                       "ETH_GAS_LIMIT_TRANSFER",
		"EthGasLimitMultiplier":                     "ETH_GAS_LIMIT_MULTIPLIER",
		"EthGasPriceDefault":                        "ETH_GAS_PRICE_DEFAULT",
		"EthHeadTrackerHistoryDepth":                "ETH_HEAD_TRACKER_HISTORY_DEPTH",
		"EthHeadTrackerMaxBufferSize":               "ETH_HEAD_TRACKER_MAX_BUFFER_SIZE",
		"EthHeadTrackerSamplingInterval":            "ETH_HEAD_TRACKER_SAMPLING_INTERVAL",
		"EthLogBackfillBatchSize":                   "ETH_LOG_BACKFILL_BATCH_SIZE",
		"EthMaxGasPriceWei":                         "ETH_MAX_GAS_PRICE_WEI",
		"EthMaxInFlightTransactions":                "ETH_MAX_IN_FLIGHT_TRANSACTIONS",
		"EthMaxQueuedTransactions":                  "ETH_MAX_QUEUED_TRANSACTIONS",
		"EthMinGasPriceWei":                         "ETH_MIN_GAS_PRICE_WEI",
		"EthNonceAutoSync":                          "ETH_NONCE_AUTO_SYNC",
		"EthRPCDefaultBatchSize":                    "ETH_RPC_DEFAULT_BATCH_SIZE",
		"EthTxReaperInterval":                       "ETH_TX_REAPER_INTERVAL",
		"EthTxReaperThreshold":                      "ETH_TX_REAPER_THRESHOLD",
		"EthTxResendAfterThreshold":                 "ETH_TX_RESEND_AFTER_THRESHOLD",
		"EthereumDisabled":                          "ETH_DISABLED",
		"EthereumHTTPURL":                           "ETH_HTTP_URL",
		"EthereumSecondaryURL":                      "ETH_SECONDARY_URL",
		"EthereumSecondaryURLs":                     "ETH_SECONDARY_URLS",
		"EthereumURL":                               "ETH_URL",
		"ExplorerAccessKey":                         "EXPLORER_ACCESS_KEY",
		"ExplorerSecret":                            "EXPLORER_SECRET",
		"ExplorerURL":                               "EXPLORER_URL",
		"FeatureCronV2":                             "FEATURE_CRON_V2",
		"FeatureExternalInitiators":                 "FEATURE_EXTERNAL_INITIATORS",
		"FeatureFluxMonitor":                        "FEATURE_FLUX_MONITOR",
		"FeatureFluxMonitorV2":                      "FEATURE_FLUX_MONITOR_V2",
		"FeatureOffchainReporting":                  "FEATURE_OFFCHAIN_REPORTING",
		"FeatureWebhookV2":                          "FEATURE_WEBHOOK_V2",
		"FlagsContractAddress":                      "FLAGS_CONTRACT_ADDRESS",
		"GasUpdaterBatchSize":                       "GAS_UPDATER_BATCH_SIZE",
		"GasUpdaterBlockDelay":                      "GAS_UPDATER_BLOCK_DELAY",
		"GasUpdaterBlockHistorySize":                "GAS_UPDATER_BLOCK_HISTORY_SIZE",
		"GasUpdaterEnabled":                         "GAS_UPDATER_ENABLED",
		"GasUpdaterTransactionPercentile":           "GAS_UPDATER_TRANSACTION_PERCENTILE",
		"GlobalLockRetryInterval":                   "GLOBAL_LOCK_RETRY_INTERVAL",
		"HTTPServerWriteTimeout":                    "HTTP_SERVER_WRITE_TIMEOUT",
		"InsecureFastScrypt":                        "INSECURE_FAST_SCRYPT",
		"JSONConsole":                               "JSON_CONSOLE",
		"JobPipelineMaxRunDuration":                 "JOB_PIPELINE_MAX_RUN_DURATION",
		"JobPipelineReaperInterval":                 "JOB_PIPELINE_REAPER_INTERVAL",
		"JobPipelineReaperThreshold":                "JOB_PIPELINE_REAPER_THRESHOLD",
		"JobPipelineResultWriteQueueDepth":          "JOB_PIPELINE_RESULT_WRITE_QUEUE_DEPTH",
		"KeeperMaximumGracePeriod":                  "KEEPER_MAXIMUM_GRACE_PERIOD",
		"KeeperMinimumRequiredConfirmations":        "KEEPER_MINIMUM_REQUIRED_CONFIRMATIONS",
		"KeeperRegistrySyncInterval":                "KEEPER_REGISTRY_SYNC_INTERVAL",
		"LinkContractAddress":                       "LINK_CONTRACT_ADDRESS",
		"LogLevel":                                  "LOG_LEVEL",
		"LogSQLMigrations":                          "LOG_SQL_MIGRATIONS",
		"LogSQLStatements":                          "LOG_SQL",
		"LogToDisk":                                 "LOG_TO_DISK",
		"MaximumServiceDuration":                    "MAXIMUM_SERVICE_DURATION",
		"MigrateDatabase":                           "MIGRATE_DATABASE",
		"MinIncomingConfirmations":                  "MIN_INCOMING_CONFIRMATIONS",
		"MinRequiredOutgoingConfirmations":          "MIN_OUTGOING_CONFIRMATIONS",
		"MinimumContractPayment":                    "MINIMUM_CONTRACT_PAYMENT",
		"MinimumRequestExpiration":                  "MINIMUM_REQUEST_EXPIRATION",
		"MinimumServiceDuration":                    "MINIMUM_SERVICE_DURATION",
		"OCRBlockchainTimeout":                      "OCR_BLOCKCHAIN_TIMEOUT",
		"OCRBootstrapCheckInterval":                 "OCR_BOOTSTRAP_CHECK_INTERVAL",
		"OCRContractConfirmations":                  "OCR_CONTRACT_CONFIRMATIONS",
		"OCRContractPollInterval":                   "OCR_CONTRACT_POLL_INTERVAL",
		"OCRContractSubscribeInterval":              "OCR_CONTRACT_SUBSCRIBE_INTERVAL",
		"OCRContractTransmitterTransmitTimeout":     "OCR_CONTRACT_TRANSMITTER_TRANSMIT_TIMEOUT",
		"OCRDHTLookupInterval":                      "OCR_DHT_LOOKUP_INTERVAL",
		"OCRDatabaseTimeout":                        "OCR_DATABASE_TIMEOUT",
		"OCRIncomingMessageBufferSize":              "OCR_INCOMING_MESSAGE_BUFFER_SIZE",
		"OCRKeyBundleID":                            "OCR_KEY_BUNDLE_ID",
		"OCRMonitoringEndpoint":                     "OCR_MONITORING_ENDPOINT",
		"OCRNewStreamTimeout":                       "OCR_NEW_STREAM_TIMEOUT",
		"OCRObservationGracePeriod":                 "OCR_OBSERVATION_GRACE_PERIOD",
		"OCRObservationTimeout":                     "OCR_OBSERVATION_TIMEOUT",
		"OCROutgoingMessageBufferSize":              "OCR_OUTGOING_MESSAGE_BUFFER_SIZE",
		"OCRTraceLogging":                           "OCR_TRACE_LOGGING",
		"OCRTransmitterAddress":                     "OCR_TRANSMITTER_ADDRESS",
		"ORMMaxIdleConns":                           "ORM_MAX_IDLE_CONNS",
		"ORMMaxOpenConns":                           "ORM_MAX_OPEN_CONNS",
		"OperatorContractAddress":                   "OPERATOR_CONTRACT_ADDRESS",
		"OptimismGasFees":                           "OPTIMISM_GAS_FEES",
		"P2PAnnounceIP":                             "P2P_ANNOUNCE_IP",
		"P2PAnnouncePort":                           "P2P_ANNOUNCE_PORT",
		"P2PBootstrapPeers":                         "P2P_BOOTSTRAP_PEERS",
		"P2PDHTAnnouncementCounterUserPrefix":       "P2P_DHT_ANNOUNCEMENT_COUNTER_USER_PREFIX",
		"P2PListenIP":                               "P2P_LISTEN_IP",
		"P2PListenPort":                             "P2P_LISTEN_PORT",
		"P2PPeerID":                                 "P2P_PEER_ID",
		"P2PPeerstoreWriteInterval":                 "P2P_PEERSTORE_WRITE_INTERVAL",
		"Port":                                      "CHAINLINK_PORT",
		"ReaperExpiration":                          "REAPER_EXPIRATION",
		"ReplayFromBlock":                           "REPLAY_FROM_BLOCK",
		"RootDir":                                   "ROOT",
		"SecureCookies":                             "SECURE_COOKIES",
		"SessionTimeout":                            "SESSION_TIMEOUT",
		"StatsPusherLogging":                        "STATS_PUSHER_LOGGING",
		"TLSCertPath":                               "TLS_CERT_PATH",
		"TLSHost":                                   "CHAINLINK_TLS_HOST",
		"TLSKeyPath":                                "TLS_KEY_PATH",
		"TLSPort":                                   "CHAINLINK_TLS_PORT",
		"TLSRedirect":                               "CHAINLINK_TLS_REDIRECT",
		"TriggerFallbackDBPollInterval":             "TRIGGER_FALLBACK_DB_POLL_INTERVAL",
		"UnAuthenticatedRateLimit":                  "UNAUTHENTICATED_RATE_LIMIT",
		"UnAuthenticatedRateLimitPeriod":            "UNAUTHENTICATED_RATE_LIMIT_PERIOD",
	}

	schemaT := reflect.TypeOf(ConfigSchema{})
	for i := 0; i < schemaT.NumField(); i++ {
		field := schemaT.FieldByIndex([]int{i})
		item, found := items[field.Name]

		//
		// ╭──╮   ╭────────────────────────────────────╮
		// │  │   │ It looks like you're trying to add │
		// @  @  ╭│ configuration variable!            │
		// ││ ││ ││                                    │
		// ││ ││ ╯╰────────────────────────────────────╯
		// │╰─╯│
		// ╰───╯
		//
		// If this test is failing, you've probably added a new configuration
		// variable, please make sure to:
		//
		// 0. Make sure that the method in config.go has a comment explaining
		//    in detail what the new config var does
		// 1. Update the changelog
		// 2. Update the ConfigPrinter found in core/store/presenters/presenters.go
		//    if you think this variable needs to be shown in the UI
		// 3. Make a PR into the documentation page if node operators might
		//    need to use this (found at https://github.com/smartcontractkit/documentation/blob/main/docs/Node%20Operators/configuration-variables.md) - don't forget to update TOC
		// 4. Add your new config variable to this test
		//

		require.True(t, found, fmt.Sprintf("New test variable: '%s', see test comment for guide on steps to follow when adding a configuration variable", field.Name))
		env := field.Tag.Get("env")
		require.Equal(t, item, env)
	}
}
