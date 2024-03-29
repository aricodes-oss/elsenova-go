import Head from 'next/head';
import Image from 'next/image';

import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';

import Nav from '@/components/Nav';

import ProfilePicture from '@/assets/profile.png';

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

          <Card style={{ width: '18rem' }} className="mx-auto my-3">
            <Card.Img variant="top" src={ProfilePicture.src} />
            <Card.Body>
              <Card.Title>💖 Aria Taylor 💖</Card.Title>
              <Card.Text>
                Musician, software engineer, Axiom Verge lover, and developer of this bot!
              </Card.Text>
              <Button variant="primary">Okay</Button>
            </Card.Body>
          </Card>
        </Container>
      </main>
    </>
  );
}
