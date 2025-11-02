# Priority demonstration: arg > env > default
# Expected output:
#   PriorityDefault: "default value" (no env or arg set)
#   PriorityEnv: "env value overwrite default" (PRIO_ENV set)
#   PriorityArg: "arg value overwrite env and default" (-prio-arg overrides PRIO_ARG)
run:
	PRIO_ENV="env value overwrite default" \
	PRIO_ARG="env value overwrite default" \
	go run main.go \
	-prio-arg="arg value overwrite env and default" \
	-password="1337" \
	-url="http://example.com" \
	-defaultWithArg="banana" \
	-active=true \
	-names="charlie,dave,eve" \
	-tags="production,api,monitoring" \
	-environments="dev|staging|prod" \
	-ports="8080:8081:9000" \
	-ids="2001,2002,2003" \
	-counters="10;20;30;40" \
	-big-numbers="9999999999,8888888888" \
	-prices="12.50,25.00,37.50" \
	-flags="true,false,true,false" \
	-allowed_users="admin,superuser,guest" \
	-broker="kafka:9092" \
	-broker-list="broker-a:9092,ssl://broker-b:9093,broker-c:9094" \
	-brokers="broker1:9092,ssl://broker2:9093,broker3:9094" \
	-std-duration=1h30m \
	-std-time=2023-12-25T10:30:00Z \
	-duration=1d2h30m \
	-datetime=2023-12-25T10:30:00Z \
	-date=2023-12-25 \
	-unixtime=1703505000 \
	-float64=13.37 \
	-int=42
