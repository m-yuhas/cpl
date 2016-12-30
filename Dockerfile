FROM swiftdocker/swift

ENV CPL_WS /opt/cpl
COPY . CPL_WS
WORKDIR CPL_WS

RUN ./build.sh install
