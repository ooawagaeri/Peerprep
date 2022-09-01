import { Request, Response } from 'express';

import ApiServer from './api_server/api_server';
import UserBFFServiceApi from './controller/user_service_bff_controller';
import UserServiceApi from './controller/user_service_controller';
import AppStorage from './storage/app_storage';

const dataStore: AppStorage = new AppStorage();

const httpPort: number = 8081;
const grpcPort: number = 4000;

const apiServer = new ApiServer(httpPort, grpcPort);
const expressApp = apiServer.getHttpServer();

expressApp.get('/', (_: Request, resp: Response) => {
  resp.status(200).send('Welcome to User Service');
});

apiServer.registerServiceRoutes(new UserServiceApi(dataStore));
apiServer.registerServiceRoutes(new UserBFFServiceApi());
apiServer.bind();