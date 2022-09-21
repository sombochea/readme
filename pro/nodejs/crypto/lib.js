const crypto = require('crypto');

const random = (size, type) => crypto.randomBytes(size).toString(type);
const randomHex = (size) => crypto.randomBytes(size).toString('hex');

module.exports = {
  random,
  randomHex,
};