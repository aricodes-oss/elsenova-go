import Container from 'react-bootstrap/Container';
import Table from 'react-bootstrap/Table';
import { useQuery, QueryClient, dehydrate, HydrationBoundary } from '@tanstack/react-query';

import Nav from '@/components/Nav';
import { vore, discord } from '@/api';

import Row from '@/components/leaderboards/Row';

function Leaderboards() {
  const { data: stats, isPending } = useQuery({
    queryKey: [...vore.queryKey, 'stats'],
    queryFn: vore.stats,
  });

  if (isPending) {
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

const LeaderboardsRoute = ({ dehydratedState }) => (
  <HydrationBoundary state={dehydratedState}>
    <Leaderboards />
  </HydrationBoundary>
);

export const getServerSideProps = async () => {
  const queryClient = new QueryClient();

  // Preload the actual stats themselves
  const stats = await queryClient.fetchQuery({
    queryKey: [...vore.queryKey, 'stats'],
    queryFn: vore.stats,
  });

  // Then fetch everyone's usernames so the table rows don't just say "loading"
  await Promise.all(
    stats.map(stat =>
      queryClient.prefetchQuery({
        queryKey: [...discord.queryKey, 'user', stat.userID],
        queryFn: () => discord.findUser(stat.userID),
      }),
    ),
  );

  return {
    props: {
      dehydratedState: dehydrate(queryClient),
    },
  };
};

export default LeaderboardsRoute;
