import { stats } from '@/api/vore';
import { useQuery } from '@tanstack/react-query';

const useStatsQuery = () => useQuery({ queryKey: ['vore', 'stats'], queryFn: stats });

export default useStatsQuery;
