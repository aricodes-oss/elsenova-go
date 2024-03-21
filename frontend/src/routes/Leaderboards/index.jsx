import Container from 'react-bootstrap/Container';
import Table from 'react-bootstrap/Table';

import Nav from '@/components/Nav';

import Row from './Row';

import { useStatsQuery } from '@/hooks/query-hooks/vore';

function Leaderboards() {
  const { data: stats, isLoading } = useStatsQuery();

  if (isLoading) {
    return 'Loading...';
  }

  return (
    <>
      <Nav />
      <Container>
        <Table striped bordered hover variant="dark" className="my-4">
          <thead>
            <tr>
              <th>User</th>
              <th>Vores</th>
            </tr>
          </thead>
          <tbody>
            {stats.map(v => (
              <Row key={v.userID} userID={v.userID} count={v.total} />
            ))}
          </tbody>
        </Table>
      </Container>
    </>
  );
}

export default Leaderboards;
