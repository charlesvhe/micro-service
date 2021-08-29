use axum::{extract::Query, handler::get, http::StatusCode, response::IntoResponse, Json, Router};
use std::net::SocketAddr;

use serde::{Deserialize, Serialize};

use consumer::consumer_client::ConsumerClient;
use provider::provider_client::ProviderClient;

pub mod consumer {
    include!("consumer.rs");
}

pub mod provider {
    include!("provider.rs");
}

#[derive(Serialize, Deserialize)]
struct Msg {
    name: String,
}

#[tokio::main]
async fn main() {
    // build our application with a route
    let app = Router::new()
        .route("/testProvider", get(testProvider))
        .route("/testConsumer", get(testConsumer));
    // run it
    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    println!("listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn testProvider(msg: Query<Msg>) -> impl IntoResponse {
    let mut client: ProviderClient<_> = ProviderClient::connect("http://localhost:50051")
        .await
        .unwrap();

    let grpcMsg = client
        .test(tonic::Request::new(provider::Msg {
            name: msg.name.clone(),
        }))
        .await
        .unwrap()
        .into_inner();

    Json(Msg { name: grpcMsg.name })
}

async fn testConsumer(msg: Query<Msg>) -> impl IntoResponse {
    let mut client: ConsumerClient<_> = ConsumerClient::connect("http://localhost:50052")
        .await
        .unwrap();

    let grpcMsg = client
        .test(tonic::Request::new(consumer::Msg {
            name: msg.name.clone(),
        }))
        .await
        .unwrap()
        .into_inner();

    Json(Msg { name: grpcMsg.name })
}
