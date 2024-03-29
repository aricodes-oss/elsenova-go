import Link from 'next/link';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

import { Github } from 'react-bootstrap-icons';

const Navigation = () => (
  <Navbar expand="lg" className="bg-body-tertiary">
    <Container>
      <Navbar.Brand as={Link} href="/">
        Elsenova
      </Navbar.Brand>

      <Navbar.Collapse>
        <Nav className="me-auto">
          <Nav.Link as={Link} href="/leaderboards">
            Leadervoreds
          </Nav.Link>
        </Nav>

        <Nav className="justify-content-end">
          <Nav.Link href="https://github.com/aricodes-oss/elsenova-go">
            <Github size={24} />
          </Nav.Link>
        </Nav>
      </Navbar.Collapse>
    </Container>
  </Navbar>
);

export default Navigation;
