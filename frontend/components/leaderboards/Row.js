import PropTypes from 'prop-types';
import { useQuery } from '@tanstack/react-query';
import { discord } from '@/api';

const Row = ({ userID, count, ...props }) => {
  const {
    data: user,
    isPending,
    ...query
  } = useQuery({
    queryKey: [...discord.queryKey, 'user', userID],
    queryFn: () => discord.findUser(userID),
  });

  return (
    <tr {...props}>
      <td>{isPending ? 'Loading...' : user.username}</td>
      <td>{count}</td>
    </tr>
  );
};

Row.propTypes = {
  userID: PropTypes.string.isRequired,
  count: PropTypes.number.isRequired,
};

export default Row;
