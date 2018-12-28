package stock.models.weather

import lombok.AllArgsConstructor
import lombok.Data

@Data
@AllArgsConstructor
data class WeatherCondition(
        val dt: Int?,
        val weatherStats: WeatherStats?,
        val weather: List<Weather?>?,
        val clouds: Clouds?,
        val wind: Wind?,
        val snow: Snow?,
        val sys: Sys?,
        val dt_txt: String?
)