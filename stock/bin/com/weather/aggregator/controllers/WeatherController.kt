package com.weather.aggregator.controllers

import com.weather.aggregator.models.weather.WeatherCondition
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.web.reactive.function.server.ServerResponse
import org.springframework.web.reactive.function.server.router

@Configuration
class WeatherController(val weatherService: WeatherService) {

    @Bean
    fun routes() = router {
        GET("/weather/{cityId}", { ServerResponse.ok().body(weatherService.getCityWeather(it.pathVariable("cityId")), WeatherCondition::class.java) })
    }
}