## Local MQTT Broker
Config File: /usr/local/etc/mosquitto/mosquitto.conf

To restart mosquitto after an upgrade:
  brew services restart mosquitto

Or, if you don't want/need a background service you can just run:
  /usr/local/opt/mosquitto/sbin/mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf
