import { LoggerEvent } from '../generated/types';

main();

function main() {
    console.log('webui started');

    const ws = new WebSocket(`ws://${window.location.host}/api/v2/events-sock`);

    ws.onopen = () => {
        console.log('[ws] connected');
    };
    ws.onmessage = (event) => {
        console.log('[ws] received message of length: ', event.data.length);
        const parsedMessage: LoggerEvent = JSON.parse(event.data);
        console.log('[ws] parsed message: ', parsedMessage);
    };
}
