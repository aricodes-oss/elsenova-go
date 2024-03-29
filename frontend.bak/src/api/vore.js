import ky from 'ky-universal';

const base = '/api/vore';

export const all = async () => await ky.get(base).json();
export const one = async id => await ky.get(`${base}/${id}`).json();
export const stats = async () => await ky.get(`${base}/stats`).json();
