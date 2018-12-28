package stock.models.weather

import lombok.AllArgsConstructor
import lombok.Data

@Data
@AllArgsConstructor
data class Coord(
        val lat: Double?,
        val lon: Double?
)