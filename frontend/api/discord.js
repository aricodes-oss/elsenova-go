import urlJoin from 'url-join';
import ky from 'ky-universal';

import APIResource from './base';
import { User } from '@/models';

class DiscordResource extends APIResource {
  resource = 'discord';
  model = User;

  findUser = async id => await ky.get(urlJoin(this.path, 'user', id)).json();
}

const discord = new DiscordResource();

export default discord;
