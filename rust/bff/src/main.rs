use axum::{extract::Query, handler::get, response::IntoResponse, Json, Router};
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
        .route("/provider", get(test_provider))
        .route("/consumer", get(test_consumer));
    // run it
    let addr = SocketAddr::from(([127, 0, 0, 1], 8888));
    println!("listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn test_provider(msg: Query<Msg>) -> impl IntoResponse {
    let mut client: ProviderClient<_> = ProviderClient::connect("http://localhost:50000")
        .await
        .unwrap();

    let grpc_msg = client
        .test(tonic::Request::new(provider::Msg {
            name: msg.name.clone(),
        }))
        .await
        .unwrap()
        .into_inner();

    Json(Msg {
        name: grpc_msg.name,
    })
}

async fn test_consumer(msg: Query<Msg>) -> impl IntoResponse {
    let mut client: ConsumerClient<_> = ConsumerClient::connect("http://localhost:60000")
        .await
        .unwrap();

    let grpc_msg = client
        .test(tonic::Request::new(provider::Msg {
            name: msg.name.clone(),
        }))
        .await
        .unwrap()
        .into_inner();

    Json(Msg {
        name: grpc_msg.name,
    })
}
