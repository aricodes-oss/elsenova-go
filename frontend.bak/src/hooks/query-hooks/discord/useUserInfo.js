import { userInfo } from '@/api/discord';
import { useQuery } from '@tanstack/react-query';

const useUserInfoQuery = id =>
  useQuery({ queryKey: ['discord', 'user', id], queryFn: () => userInfo(id) });

export default useUserInfoQuery;
