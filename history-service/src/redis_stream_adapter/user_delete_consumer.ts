import { commandOptions, RedisClientType } from 'redis';
import Logger from '../utils/logger';
import { IAttemptStore } from '../storage/storage';

const STREAMS_KEY = 'stream-delete-user';

const CONSUMER_GROUP = 'stream-delete-user-consumer';

const CONSUMER_NAME = 'stream-consumer-question-service';

class UserDeleteConsumer implements IStreamConsumer {
  redis: RedisClientType;

  storage: IAttemptStore;

  constructor(redis: RedisClientType, storage: IAttemptStore) {
    this.redis = redis;
    this.storage = storage;
  }

  async runConsumer() {
    try {
      await this.redis.xGroupCreate(STREAMS_KEY, CONSUMER_GROUP, '0', {
        MKSTREAM: true,
      });
      Logger.info('Created Redis User consumer group.');
    } catch (e) {
      Logger.info('Consumer Redis User group already exists, skipped creation.');
    }

    while (true) {
      try {
        // eslint-disable-next-line no-await-in-loop
        const response = await this.redis.xReadGroup(
          // @ts-ignore
          commandOptions({ isolated: true }),
          CONSUMER_GROUP,
          CONSUMER_NAME,
          [
            {
              key: STREAMS_KEY,
              id: '>',
            },
          ],
          {
            COUNT: 1,
            BLOCK: 5000,
          },
        );
        if (response) {
          const { userId } = response[0].messages[0].message;
          this.storage.removeHistoryOwner(Number(userId));

          const entryId = response[0].messages[0].id;
          this.redis.xAck(STREAMS_KEY, CONSUMER_GROUP, entryId);
        }
      } catch (err) {
        Logger.error('Redis User consumer failed');
      }
    }
  }
}

function createUserDeleteConsumer(redis: RedisClientType, storage: IAttemptStore): IStreamConsumer {
  return new UserDeleteConsumer(redis, storage);
}

export default createUserDeleteConsumer;
