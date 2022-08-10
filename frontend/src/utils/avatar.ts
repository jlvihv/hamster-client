import { createAvatar } from '@dicebear/avatars';
import * as style from '@dicebear/avatars-identicon-sprites';

// - seed: a indetify string to generate avatar
// - options: https://avatars.dicebear.com/docs/options
export function createSvgAvatar(seed: string, options: Recordable = {}) {
  return createAvatar(style, { seed, dataUri: true, ...options });
}
