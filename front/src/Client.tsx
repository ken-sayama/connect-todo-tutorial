import {FC, ReactNode} from "react";
import {createConnectTransport} from "@connectrpc/connect-web";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {TransportProvider} from "@bufbuild/connect-query";

type Props = {
    baseUrl: string;
    children: ReactNode;
}

export const Client: FC<Props> = ({ baseUrl, children }) => {
    const transport = createConnectTransport({
        baseUrl,
        interceptors: [],
    });
    const client = new QueryClient();

    return (
        <TransportProvider transport={transport}>
            <QueryClientProvider client={client}>
                {children}
            </QueryClientProvider>
        </TransportProvider>
    )
}