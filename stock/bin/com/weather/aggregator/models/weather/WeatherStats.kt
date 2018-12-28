package com.weather.aggregator.models.weather

import lombok.AllArgsConstructor
import lombok.Data

@Data
@AllArgsConstructor
data class WeatherStats(
        val temp: Double?,
        val temp_min: Double?,
        val temp_max: Double?,
        val pressure: Double?,
        val sea_level: Double?,
        val grnd_level: Double?,
        val humidity: Int?,
        val temp_kf: Double?
)