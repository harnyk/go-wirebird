import { useCallback } from 'react';
import { useEffectOnce, useList, useSet } from 'react-use';
import { LoggerEvent } from '../../../generated/types';

const useLookups = () => {
    //domains lookup
    const [domains, { add: addDomain }] = useSet<string>();
    //pid lookup
    const [pids, { add: addPid }] = useSet<number | null>();
    //method lookup
    const [methods, { add: addMethod }] = useSet<string>();
    //status code lookup
    const [statusCodes, { add: addStatusCode }] = useSet<number | null>();

    const indexEvent = useCallback(
        (event: LoggerEvent) => {
            addDomain(new URL(event.request.url).hostname);
            addPid(event.process?.pid ?? null);
            addMethod(event.request.method);
            addStatusCode(event.response?.statusCode ?? null);
        },
        [addDomain, addPid, addMethod, addStatusCode]
    );

    return {
        lookups: { domains, pids, methods, statusCodes },
        indexEvent,
    };
};

export const useEventStream = (url: string) => {
    const [events, { push }] = useList<LoggerEvent>([]);
    const { lookups, indexEvent } = useLookups();

    useEffectOnce(() => {
        const ws = new WebSocket(url);

        ws.onopen = () => {
            console.log('[ws] connected');
        };
        ws.onmessage = (event) => {
            console.log('[ws] received message of length: ', event.data.length);
            const parsedMessage: LoggerEvent = JSON.parse(event.data);
            console.log('[ws] parsed message: ', parsedMessage);
            push(parsedMessage);
            indexEvent(parsedMessage);
        };

        return () => {
            ws.close();
        };
    });

    return { events, lookups };
};
