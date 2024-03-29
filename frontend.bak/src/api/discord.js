import ky from 'ky-universal';

const base = '/api/discord';

export const userInfo = async id => await ky.get(`${base}/user/${id}`).json();
