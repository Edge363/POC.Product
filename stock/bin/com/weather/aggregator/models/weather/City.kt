package com.weather.aggregator.models.weather

import lombok.AllArgsConstructor
import lombok.Data

@Data
@AllArgsConstructor
data class City(
        val id: Int?,
        val name: String?,
        val coord: Coord?,
        val country: String?
)