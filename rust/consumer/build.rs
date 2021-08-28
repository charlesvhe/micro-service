fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::compile_protos("proto/consumer.proto")?;
    tonic_build::compile_protos("proto/provider.proto")?;
    Ok(())
}