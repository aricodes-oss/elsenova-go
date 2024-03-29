import Head from 'next/head';

import Container from 'react-bootstrap/Container';

import Nav from '@/components/Nav';

export default function Home() {
  return (
    <>
      <Head>
        <title>Elsenova</title>
        <meta name="description" content="A bot for the Axiom Verge discord" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Nav />

      <main>
        <Container>
          <h1 className="mt-5">Elsenova</h1>
          <p className="my-2">
            Hey! We uh. We don't have a whole lot here right now. Stay tuned though!
          </p>
        </Container>
      </main>
    </>
  );
}
