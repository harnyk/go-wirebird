import Fuse from 'fuse.js';
import { useCallback, useMemo, useRef } from 'react';
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

interface FilterParams {
    domain?: string;
    pid?: number;
    method?: string;
    statusCode?: number;
}

const useSearchIndex = (filter: FilterParams) => {
    const index = useMemo(
        () =>
            new Fuse<LoggerEvent>([], {
                keys: [
                    'request.url',
                    'request.method',
                    'response.statusCode',
                    'process.pid',
                ],
            }),
        []
    );

    const versionRef = useRef(0);

    const indexEvent = useCallback((event: LoggerEvent) => {
        index.add(event);
        versionRef.current++;
    }, []);

    const searchResults = useMemo(() => {
        const { domain, pid, method, statusCode } = filter;

        const expressions: Fuse.Expression[] = [];
        if (domain) {
            expressions.push({ 'request.url': domain });
        }
        if (pid) {
            expressions.push({ 'process.pid': pid.toString() });
        }
        if (method) {
            expressions.push({ 'request.method': method });
        }
        if (statusCode) {
            expressions.push({ 'response.statusCode': statusCode.toString() });
        }

        const results = index.search({ $and: expressions });
        return results.map((r) => r.item);
    }, [
        filter.domain,
        filter.pid,
        filter.method,
        filter.statusCode,
        versionRef.current,
    ]);

    return {
        indexEvent,
        searchResults,
    };
};

export const useEventStream = (url: string, filter: FilterParams) => {
    const [events, { push }] = useList<LoggerEvent>([]);
    const { lookups, indexEvent } = useLookups();
    const { indexEvent: indexSearch, searchResults } = useSearchIndex(filter);

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
            indexSearch(parsedMessage);
        };

        return () => {
            ws.close();
        };
    });

    return { events, searchResults, lookups };
};
