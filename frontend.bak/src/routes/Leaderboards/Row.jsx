import PropTypes from 'prop-types';
import { useUserInfoQuery } from '@/hooks/query-hooks/discord';

const Row = ({ userID, count, ...props }) => {
  const { data: user, isLoading } = useUserInfoQuery(userID);

  return (
    <tr {...props}>
      <td>{isLoading ? 'Loading...' : user.username}</td>
      <td>{count}</td>
    </tr>
  );
};

Row.propTypes = {
  userID: PropTypes.string.isRequired,
  count: PropTypes.number.isRequired,
};

export default Row;
