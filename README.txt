docker-compose up --build -d


Request examples:

localhost:8080/devices POST

{
    "device_id": 23,
    "plate_number": "20",
    "timestamp": "11-02-2022 15:49:56"
}

{
    "device_id": 23,
    "temperature": "20",
    "unit":"celsius",
    "timestamp": "11-02-2022 15:49:56"
}


{
    "device_id": 23,
    "velocity": "20",
    "unit":"km/h",
    "timestamp": "11-02-2022 15:49:56"
}

localhost:8080/devices GET

device_types = traffic,velocity,temperature
{
    "device_type": "traffic",
    "start_date": "10-02-2022 07:49:56",
    "end_date": "11-02-2022 19:49:56",
    "limit": 10
}