use tonic::{transport::Server, Request, Response, Status};

use consumer::consumer_server::{Consumer, ConsumerServer};
use provider::Msg;

use provider::provider_client::ProviderClient;

pub mod consumer {
    include!("consumer.rs");
}

pub mod provider {
    include!("provider.rs");
}

#[derive(Debug, Default)]
pub struct MyConsumer {}

#[tonic::async_trait]
impl Consumer for MyConsumer {
    async fn test(&self, request: Request<Msg>) -> Result<Response<Msg>, Status> {
        println!("Got a request: {:?}", request);

        let mut client: ProviderClient<_> = ProviderClient::connect("http://localhost:50000")
            .await
            .unwrap();

        let msg = tonic::Request::new(provider::Msg {
            name: request.into_inner().name,
        });

        let reply = provider::Msg {
            name: format!(
                "Consumer Hello {}!",
                client.test(msg).await.unwrap().into_inner().name
            )
            .into(),
            // name: format!("Consumer Hello {}!", request.into_inner().name).into(),
        };

        Ok(Response::new(reply))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:60000".parse()?;
    let consumer = MyConsumer::default();

    Server::builder()
        .add_service(ConsumerServer::new(consumer))
        .serve(addr)
        .await?;

    Ok(())
}
