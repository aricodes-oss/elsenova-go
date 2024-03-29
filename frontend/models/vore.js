import Resource from './base';
import { createModelSchema, identifier, primitive } from 'serializr';

class Vore extends Resource {
  id = 0;
  userID = 0;
}

createModelSchema(Vore, {
  id: identifier(),
  userID: primitive(),
});

export default Vore;
