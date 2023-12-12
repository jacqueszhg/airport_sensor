# Golang Project - Airport Sensor Simulation

## Introduction

This project aims to simulate airport sensors, specifically those measuring wind speed, temperature, and altitude. Simulated data is stored in an InfluxDB NoSQL database and saved in Excel files. The backend is developed in Go, utilizing MQTT for communication between sensors and the server, InfluxDB for data storage, and a Vue.js frontend for the user interface.

## Project Structure

### Building the Project

To compile the project, run the `script\build.bat` script. You can configure the installation path and the airport's IATA code in this script.

```batch
SET INSTALL_PATH="C:\Users\lucas\airport" // Folder to generate executables
SET AIRPORT="NTE" // Airport's IATA code
```

### Launching Services

To start the MQTT Mosquitto broker, database connection, logs, and API, execute the `runservices.bat` script located in the project folder.

### Launching Sensors

To start the sensors for an airport, execute the `runsensors.bat` script in the airport's folder, named with its IATA code. Each sensor has its own `config.yml` file.

```yaml
mqtt:
  protocol: 'tcp'
  url: 'localhost'
  port: '1883'

sensor:
  id: 1 // Unique identifier for each sensor
  airport: 'NTE' // Airport's IATA code
  frequency: 10 // Data sending frequency in seconds
  QOSLevel: 1 // Sensor data sending quality
  altitudeAirport: 100 // Airport altitude for simulation
```

### User Interface (UI) in `website`

To launch the UI, run the following commands in the `website` folder:

```bash
npm install # Install dependencies
npm run dev # Launch the project
```

### Launching Services Independently

#### Mosquitto (MQTT Broker)

Start the MQTT broker Mosquitto using the following command in a terminal:

```bash
mosquitto -v
```

#### Publishing, Subscribing to Logs and Database, and API

Launch the executables one by one located in the folders:

-   sub
-   api
-   IATA/sensors/pressure
-   IATA/sensors/temperature
-   IATA/sensors/wind

#### API

The Swagger documentation for the API is available at: `http://localhost:8080/swagger/index.html`

If you don't have the `swag` command, install it using the following command:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

To update the Swagger documentation, execute the command:

```bash
swag init
```
