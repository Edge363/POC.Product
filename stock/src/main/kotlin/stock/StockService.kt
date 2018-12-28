package stock

import stock.models.weather.WeatherCondition
import stock.models.weather.WeatherConditions
import org.springframework.stereotype.Service
import org.springframework.web.client.RestTemplate
import reactor.core.publisher.Mono
import reactor.core.publisher.toMono

@Service
class StockService {

    fun getCityWeather(cityId: String): Mono<WeatherCondition> {
        val response = RestTemplate()
                       .getForObject("http://api.openweathermap.org/data/2.5/forecast?id=$cityId&APPID=fecbe58a9bad9eb15822ae9d7cbe0266&cnt=1", WeatherConditions::class.java)
                       .list!![0]
       return when(response) {
            null -> WeatherCondition(null,null,null,null,null,null,null,null).toMono()
            else ->response.toMono()
        }
    }

    fun hello(): Mono<String> {
            return "hello".toMono()
    }

}