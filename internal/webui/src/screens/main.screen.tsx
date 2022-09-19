import React, { FC } from 'react';
import { useEventStream } from '../hooks/useEventStream';

export const MainScreen: FC = () => {
    const { events, lookups } = useEventStream(
        `ws://${window.location.host}/api/v2/events-sock`
    );

    return (
        <div>
            <h1>Main Screen</h1>
            <ul>
                {events.map((event) => (
                    <li key={event.request.id}>{event.request.id}</li>
                ))}
            </ul>
            <pre>{JSON.stringify([...lookups.domains], null, 2)}</pre>
        </div>
    );
};
