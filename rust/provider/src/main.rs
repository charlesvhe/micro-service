use tonic::{transport::Server, Request, Response, Status};

use provider::provider_server::{Provider, ProviderServer};
use provider::Msg;

pub mod provider {
    include!("provider.rs");
}

#[derive(Debug, Default)]
pub struct MyProvider {}

#[tonic::async_trait]
impl Provider for MyProvider {
    async fn test(&self, request: Request<Msg>) -> Result<Response<Msg>, Status> {
        println!("Got a request: {:?}", request);

        let reply = provider::Msg {
            name: format!("Hello {}!", request.into_inner().name).into(),
        };

        Ok(Response::new(reply))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:50000".parse()?;
    let provider = MyProvider::default();

    Server::builder()
        .add_service(ProviderServer::new(provider))
        .serve(addr)
        .await?;

    Ok(())
}
