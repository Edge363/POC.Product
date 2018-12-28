package stock

import stock.models.weather.WeatherCondition
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.web.reactive.function.server.ServerResponse
import org.springframework.web.reactive.function.server.router

@Configuration
class StockController(val stockService: StockService) {
    @Bean
    fun routes() = router {
        GET("/weather/{cityId}", { ServerResponse.ok().body(stockService.getCityWeather(it.pathVariable("cityId")), WeatherCondition::class.java) })
        GET("/", { ServerResponse.ok().body(stockService.hello(),String::class.java) })
    }
}