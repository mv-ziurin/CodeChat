FROM microsoft/dotnet:2.1-sdk
WORKDIR /app

COPY . ./
RUN dotnet restore UnitTest/UnitTest.csproj
RUN dotnet publish -c Release -o out

WORKDIR UnitTest/out

ENTRYPOINT ["dotnet", "vstest", "UnitTest.dll"]
