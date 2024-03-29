import urlJoin from 'url-join';
import ky from 'ky-universal';
import APIResource from './base';
import { Vore } from '@/models';

class VoreResource extends APIResource {
  resource = 'vore';
  model = Vore;

  all = async () => await this.ky.get(this.path).json();
  find = async id => await this.ky.get(urlJoin(this.path, id)).json();
  // Using raw ky here since this doesn't fit the model schema
  stats = async () => await ky.get(urlJoin(this.path, 'stats')).json();
}

const vore = new VoreResource();

export default vore;
