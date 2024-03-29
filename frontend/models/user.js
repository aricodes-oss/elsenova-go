import Resource from './base';
import { createModelSchema, identifier, primitive } from 'serializr';

class User extends Resource {
  id = 0;
  email = '';
  username = '';
  avatar = '';
  locale = '';
  discriminator = '';
  verified = false;
  banner = '';
}

createModelSchema(User, {
  id: identifier(),
  email: primitive(),
  username: primitive(),
  avatar: primitive(),
  locale: primitive(),
  discriminator: primitive(),
  verified: primitive(),
  banner: primitive(),
});

export default User;
