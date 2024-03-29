import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

import { Github } from 'react-bootstrap-icons';
import { LinkContainer } from 'react-router-bootstrap';

const Navigation = () => (
  <Navbar expand="lg" className="bg-body-tertiary">
    <Container>
      <LinkContainer to="/">
        <Navbar.Brand>Elsenova</Navbar.Brand>
      </LinkContainer>

      <Navbar.Collapse>
        <Nav className="me-auto">
          <LinkContainer to="/leaderboards">
            <Nav.Link>Leadervoreds</Nav.Link>
          </LinkContainer>
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
