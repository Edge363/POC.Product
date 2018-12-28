package com.weather.aggregator.models.weather

import lombok.AllArgsConstructor
import lombok.Data

@Data
@AllArgsConstructor
data class Sys(
        val pod: String?
)