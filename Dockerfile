#############
# Base
#############

FROM mcr.microsoft.com/dotnet/sdk:6.0 as base
RUN apt update && apt upgrade -y
RUN apt install todotxt-cli
WORKDIR /mnt/src
EXPOSE 5000


#############
# Development
#############

FROM base as dev

# requires a volume mount
ENTRYPOINT ["dotnet", "watch", "run"]


#############
# Build
#############

FROM base as build
COPY src .
RUN dotnet restore && dotnet build
RUN dotnet publish ChaosTasks.csproj --runtime alpine-x64 --self-contained true -c Release -o /dist


#####
# UAT
#####

FROM mcr.microsoft.com/dotnet/runtime:6.0-alpine3.15 as uat
COPY --from=build /dist /var/www
ENTRYPOINT ["dotnet", "/var/www/ChaosTasks.dll"]


############
# Production
############

FROM mcr.microsoft.com/dotnet/runtime:6.0-alpine3.15 as prod
COPY --from=build /dist /var/www
ENTRYPOINT ["dotnet", "/var/www/ChaosTasks.dll"]
