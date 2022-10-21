import { RedisClientType } from 'redis';

type MatchResult = {
  matched: boolean;
  difficulty: number;
  roomId: string;
};

interface IRedisMatchingAdapter {
  pushStream(username: string, difficulties: number[]): Promise<string | undefined>;
  lockIfUnset(username: string): Promise<boolean>;
  getUserLock(username: string): Promise<MatchResult | null>;
  deleteUserLock(username: string): Promise<boolean>;
}

const MATCHMAKER_QUEUE_KEY = 'matchmaker-stream';
const MATCHMAKER_LOCK_EXPIRY = 32; // Add 2 seconds as a buffer to prevent requeue
const getMatchmakerUserKey = (username: string) => `matchmaker-${username}`;

class RedisMatchingAdapter implements IRedisMatchingAdapter {
  redisClient: RedisClientType;

  constructor(redisClient: RedisClientType) {
    this.redisClient = redisClient;
  }

  async pushStream(username: string, difficulties: number[]): Promise<string | undefined> {
    const queueId = await this.redisClient.xAdd(
      MATCHMAKER_QUEUE_KEY,
      '*',
      RedisMatchingAdapter.createQueueItem(username, JSON.stringify(difficulties)),
    );
    if (queueId === '') {
      return undefined;
    }
    return queueId;
  }

  async getUserLock(username: string): Promise<MatchResult | null> {
    const key = getMatchmakerUserKey(username);
    const result = await this.redisClient.get(key);
    if (result === null) {
      return result;
    }

    if (!result.includes(';;')) {
      return {
        matched: false,
        difficulty: 0,
        roomId: '',
      };
    }

    const tokenParts = result.split(';;');
    if (tokenParts.length < 2) {
      return null;
    }

    const difficulty = parseInt(tokenParts[0], 10);
    if (Number.isNaN(difficulty)) {
      return null;
    }

    return {
      matched: true,
      difficulty,
      roomId: tokenParts[1],
    };
  }

  async deleteUserLock(username: string): Promise<boolean> {
    const key = getMatchmakerUserKey(username);
    const result = await this.redisClient.del(key);
    if (result === 0) {
      return false;
    }
    return true;
  }

  async lockIfUnset(username: string): Promise<boolean> {
    const key = getMatchmakerUserKey(username);
    const result = await this.redisClient.set(key, '', {
      NX: true,
      GET: true,
      EX: MATCHMAKER_LOCK_EXPIRY,
    });

    if (result !== null) {
      return false;
    }
    return true;
  }

  static createQueueItem(username: string, difficulty: string): Record<string, string> {
    return {
      user: username,
      diff: difficulty,
    };
  }
}

function createRedisMatchingAdapter(redisClient: RedisClientType): IRedisMatchingAdapter {
  return new RedisMatchingAdapter(redisClient);
}

export {
  IRedisMatchingAdapter,
  createRedisMatchingAdapter,
  MatchResult,
};
