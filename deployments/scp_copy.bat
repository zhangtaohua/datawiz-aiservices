setlocal

rem vars
set USER=root
set PASSWORD=123456
set HOST=192.168.3.237
set REMOTE_DIR=/home/datawiz/datawiz-ai
set SOURCE_DIR_BUILD=D:\Work\project\golang\src\datawiz-aiservices\build
set SOURCE_DIR_DEPLOY=D:\Work\project\golang\src\datawiz-aiservices\deployments
set SOURCE_DIR_env=D:\Work\project\golang\src\datawiz-aiservices\.env
set SOURCE_DIR_MOUNT=D:\Work\project\golang\src\datawiz-aiservices\deployments\mount_and_copy.sh

rem scp
scp -r %SOURCE_DIR_BUILD% %USER%@%HOST%:%REMOTE_DIR%
scp -r %SOURCE_DIR_DEPLOY% %USER%@%HOST%:%REMOTE_DIR%
scp -r %SOURCE_DIR_env% %USER%@%HOST%:%REMOTE_DIR%
rem scp -r %SOURCE_DIR_MOUNT% %USER%@%HOST%:%REMOTE_DIR%

rem pscp
rem pscp -r -pw %PASSWORD% %SOURCE_DIR% %USER%@%HOST%:%REMOTE_DIR%

endlocal
echo Files copied successfully.
pause
