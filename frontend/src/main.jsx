import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import Home from '@/routes/Home';
import Leaderboards from '@/routes/Leaderboards';

import 'bootstrap/dist/css/bootstrap.min.css';

const router = createBrowserRouter([
  { path: '/', element: <Home /> },
  { path: '/leaderboards', element: <Leaderboards /> },
]);

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>,
);
