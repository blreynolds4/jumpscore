import axios from 'axios';

const getEvents = () => {
    const response = axios.get('http://localhost:8080/api/events',
    {
        headers: {
            "Access-Control-Allow-Origin": "*",
        }
    });

    console.log(response);
    return response
};

export default getEvents;