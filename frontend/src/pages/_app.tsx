import type { AppProps } from "next/app";
import { ApolloClient, InMemoryCache, createHttpLink } from "@apollo/client";
import { ApolloProvider } from "@apollo/client";
import { NextUIProvider } from '@nextui-org/react';

function MyApp({ Component, pageProps }: AppProps) {
  const link = createHttpLink({
    uri: "http://localhost:8080/query",
    credentials: "include",
  });
  const client = new ApolloClient({
    cache: new InMemoryCache(),
    link: link,
  });

  return (
    <ApolloProvider client={client}>
      <NextUIProvider>
        <Component {...pageProps} />
      </NextUIProvider>
    </ApolloProvider>
  );
}

export default MyApp;