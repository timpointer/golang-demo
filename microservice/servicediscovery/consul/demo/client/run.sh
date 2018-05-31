.\consul.exe agent -dev -config-dir=D:\workspace\projects\bin\consul\consul.d

consul agent -server -bootstrap-expect=0 \
    -data-dir=D:\workspace\projects\bin\consul\tmp -node=agent-one -bind=172.0.0.1 \
    -enable-script-checks=true -config-dir=D:\workspace\projects\bin\consul\consul.d

consul agent -server -bootstrap-expect=1 \
    -data-dir=D:\workspace\projects\bin\consul\tmp -node=agent-one -bind=172.0.0.1 \
    -enable-script-checks=true -config-dir=D:\workspace\projects\bin\consul\consul.d