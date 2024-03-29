import urlJoin from 'url-join';
import APIResource from './base';
import { User } from '@/models';

class DiscordResource extends APIResource {
  resource = 'discord';
  model = User;

  findUser = async id => await this.ky.get(urlJoin(this.path, 'user', id)).json();
}

const discord = new DiscordResource();

export default discord;
