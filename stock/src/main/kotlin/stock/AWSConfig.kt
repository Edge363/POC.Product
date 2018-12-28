// package stock

// import com.amazonaws.services.rds.AmazonRDSAsync
// import com.amazonaws.services.rds.AmazonRDSAsyncClientBuilder
// import org.springframework.beans.factory.annotation.Value
// import org.springframework.context.annotation.Bean
// import org.springframework.context.annotation.Configuration

// @Configuration
// class AWSConfig {

//     @Bean
//     fun getRDS(@Value("RDS_ADDRESS") rdsdns: String): AmazonRDSAsync? {
//         return AmazonRDSAsyncClientBuilder.defaultClient()
//     }

// }