package com.weather.aggregator.models.weather

import lombok.AllArgsConstructor
import lombok.Data

@Data
@AllArgsConstructor
data class Clouds(
        val all: Int?
)