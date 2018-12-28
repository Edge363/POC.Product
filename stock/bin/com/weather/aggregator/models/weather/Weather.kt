package com.weather.aggregator.models.weather

import lombok.AllArgsConstructor
import lombok.Data

@Data
@AllArgsConstructor
data class Weather(
        val id: Int?,
        val main: String?,
        val description: String?,
        val icon: String?
)