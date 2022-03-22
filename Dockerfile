#############
# Base
#############

FROM mcr.microsoft.com/dotnet/sdk:6.0 as base
RUN apt update && apt upgrade -y
RUN apt install todotxt-cli
WORKDIR /mnt/src


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
RUN dotnet publish App.csproj -c Release -o /dist


#####
# UAT
#####

FROM mcr.microsoft.com/dotnet/runtime:6.0 as uat
COPY --from=build /dist /var/www
ENTRYPOINT ["dotnet", "/var/www/App.dll"]


############
# Production
############

FROM mcr.microsoft.com/dotnet/runtime:6.0 as prod
COPY --from=build /dist /var/www
ENTRYPOINT ["dotnet", "/var/www/App.dll"]
