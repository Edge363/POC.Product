package stock.models.weather

import lombok.AllArgsConstructor
import lombok.Data

@Data
@AllArgsConstructor
data class Wind(
        val speed: Double?,
        val deg: Double?
)